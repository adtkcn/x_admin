import { toast } from "@/utils/utils";

/**
 * 调起地图
 * @param {number} lat: 纬度gcj02
 * @param {number} lon: 经度
 * @param {string} title: 地址
 */
export function openMap(lat, lon, title) {
  // #ifdef APP-PLUS
  openGaoDe(lat, lon, title).catch(() => {
    openBaidu(lat, lon, title).catch(() => {
      toast("调起地图失败,请安装百度或高德地图");
      setTimeout(() => {
        openGaoDeWeb(lat, lon, title);
      }, 1000);
    });
  });
  // #endif
  // #ifdef H5
  openGaoDeWeb(lat, lon, title);
  // #endif
}
/**
 * 调起高德地图
 * @param {number} lat: 纬度gcj02
 * @param {number} lon: 经度
 * @param {string} title: 地址
 */
function openGaoDe(lat, lon, title) {
  return new Promise((resolve, reject) => {
    let gaodeURl = `androidamap://viewMap?sourceApplication=售货机维护&poiname=${title}&lat=${lat}&lon=${lon}&dev=0`;

    //#ifdef APP-IOS
    gaodeURl = `iosamap://navi?sourceApplication=售货机维护&poiname=${title}&lat=${lat}&lon=${lon}&dev=0`;
    //#endif
    console.log("gaodeURl", gaodeURl);
    plus.runtime.openURL(encodeURI(gaodeURl), function (res) {
      // 打开指定URL地址失败时回调，并返回失败信息。
      console.log("打开指定URL地址失败时回调", res);
      //   toast("调起高德地图失败");
      reject();
    });
    setTimeout(() => {
      resolve(true);
    }, 1000);
  });
}
/**
 * 调起百度地图
 * @param {number} lat: 纬度gcj02
 * @param {number} lon: 经度
 * @param {string} title: 地址
 */
function openBaidu(lat, lon, title) {
  return new Promise((resolve, reject) => {
    let baiduUrl = `baidumap://map/marker?location=${lat},${lon}&title=${title}&coord_type=gcj02&traffic=on&src=uni.UNIF9B953B`;
    console.log("baiduUrl", baiduUrl);
    plus.runtime.openURL(encodeURI(baiduUrl), function (res) {
      // 打开指定URL地址失败时回调，并返回失败信息。
      console.log("baiduUrl打开指定URL地址失败时回调", res);
      //   toast("调起百度地图失败");
      reject();
    });
    setTimeout(() => {
      resolve(true);
    }, 1000);
  });
}
/**
 * 网页调起高德地图
 * @param {number} lat: 纬度gcj02
 * @param {number} lon: 经度
 * @param {string} title: 地址
 */
function openGaoDeWeb(lat, lon, title) {
  let gaodeURl = `https://uri.amap.com/marker?position=${lon},${lat}&name=${title}&src=uni.UNIF9B953B&coordinate=gaode&callnative=1`;
  console.log("gaodeURl", gaodeURl);

  // #ifdef APP-PLUS
  plus.runtime.openWeb(gaodeURl);
  // #endif
  // #ifdef H5
  window.open(gaodeURl);
  // #endif
}
