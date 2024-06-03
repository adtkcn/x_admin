import { request } from "@/utils/request";
import { reactive, toRaw } from "vue";

interface Options {
  [propName: string]: string
}

export function useDictOptions<T = any>(options: Options) {
  const optionsData: any = reactive({});
  const optionsKey = Object.keys(options);
  
  const apiLists = optionsKey.map((key) => {
    const path = options[key];
    optionsData[key] = [];
    return () =>request({
        url: path,
        method: "GET",
      });
  });

  const refresh = async () => {
    const res = await Promise.allSettled<Promise<any>>(
      apiLists.map((api) => api())
    );
    console.log(res);

    res.forEach((item, index) => {
      const key = optionsKey[index];
      if (item.status == "fulfilled") {
        const data = item.value;
        optionsData[key] = data.data;
      }
    });
  };
  refresh();
  return {
    optionsData: optionsData as T,
    refresh,
  };
}

export function useDictData<T = any>(dict: string[]) {
  const options: Options = {};
  for (const type of dict) {
    options[type] =  `/setting/dict/data/all?dictType=${type}`
  }
  const { optionsData } = useDictOptions<T>(options);
  console.log('optionsData',optionsData);
  
  return {
    dictData: optionsData as T,
  };
}

// export function useAllList<T = any>(options: Options) {}
