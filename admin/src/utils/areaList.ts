// 下载范围区域列表
export interface AreaNode {
  areaCode: string | number
  areaName: string
  provinceCode?: string
  cityCode?: string
  fetchLoad: () => Promise<any>
  children?: AreaNode[]
}

const modules: Record<string, () => Promise<any>> = import.meta.glob('../assets/geojson/**/*.json')

function formate(): AreaNode[] {
  const pre = '../assets/geojson/'
  const ext = '.json'
  const provinceId = 'province'
  const cityId = 'city'
  const contryList: AreaNode[] = []
  const provinceList: AreaNode[] = []
  const cityMap: Record<string, AreaNode[]> = {}
  for (const path in modules) {
    const params = path.substring(pre.length, path.length - ext.length).split('/')
    if (params.length === 1) {
      contryList.push({
        areaCode: -1,
        areaName: '中国',
        fetchLoad: modules[path]
      })
      continue
    }
    const [folderName, fileName] = params
    const [areaCode, areaName] = fileName.split('-')
    const provinceCode = areaCode.substring(0, 2) // 省份代码
    const cityCode = areaCode.substring(2, 4) // 城市代码
    // const countyCode = areaCode.substring(4, 6); // 区/乡镇代码
    if (folderName === provinceId) {
      provinceList.push({
        areaCode,
        provinceCode,
        cityCode,
        areaName,
        fetchLoad: modules[path]
      })
    } else if (folderName === cityId) {
      if (!cityMap[provinceCode]) cityMap[provinceCode] = []
      cityMap[provinceCode].push({
        areaCode,
        provinceCode,
        cityCode,
        areaName,
        fetchLoad: modules[path]
      })
    }
  }
  for (let index = 0; index < provinceList.length; index++) {
    const province = provinceList[index]
    if (province.provinceCode) province.children = cityMap[province.provinceCode]
  }
  contryList[0].children = provinceList
  return contryList
}

const areaList: AreaNode[] = formate()

export function getAreaList(): AreaNode[] {
  return areaList
}
