import { reactive, toRaw } from "vue";
import { toast } from "@/utils/utils";

// 分页钩子函数
interface Options {
  pageNo?: number;
  pageSize?: number;
  loadingText?: string;
  loadmoreText?: string;
  nomoreText?: string;

  fetchFun: (_arg: any) => Promise<any>;
  params?: Record<any, any>;
  firstStatus?: string; //loadmore - 加载前，loading - 加载中，nomore - 没有数据
}

export function usePaging(options: Options) {
  const {
    pageNo = 1,
    pageSize = 30,
    loadingText = "努力加载中",
    loadmoreText = "继续加载",
    nomoreText = "没有数据了",
    fetchFun,
    params = {},
    firstStatus = "loadmore",
  } = options;
  // 记录分页初始参数
  const paramsInit: Record<any, any> = Object.assign({}, toRaw(params));
  // 分页数据
  const pager = reactive({
    pageNo,
    pageSize,
    loadingText,
    loadmoreText,
    nomoreText,
    loading: firstStatus,
    count: 0,
    lists: [] as any[],
  });
  // 请求分页接口
  const getLists = () => {
    pager.loading = "loading";
    return fetchFun({
      pageNo: pager.pageNo,
      pageSize: pager.pageSize,
      ...params,
    })
      .then((res: any) => {
        uni.stopPullDownRefresh();

        if (res.code == 200) {
          pager.count = res?.data?.count;
          var list = res?.data?.lists || [];

          pager.lists.push(...list);

          if (res?.data?.lists?.length < pager.pageSize) {
            pager.loading = "nomore";
          } else {
            pager.loading = "loadmore";
          }
        }else{
          pager.loading = "loadmore";
        }
        return Promise.resolve(res);
      })
      .catch((err: any) => {
        uni.stopPullDownRefresh();
        toast(err.message);
        pager.loading = "nomore";
        return Promise.reject(err);
      });
  };
  // 下一页
  const NextPage = () => {
    if (pager.loading === "nomore") {
      toast("没有数据了");
      return;
    }
    pager.pageNo += 1;
    getLists();
  };

  // 重置为第一页
  const resetPage = () => {
    pager.pageNo = 1;
    pager.lists = [];
    getLists();
  };
  // 重置参数
  const resetParams = () => {
    Object.keys(paramsInit).forEach((item) => {
      params[item] = paramsInit[item];
    });
    pager.lists = [];
    getLists();
  };
  return {
    pager,
    getLists,
    NextPage,
    resetParams,
    resetPage,
  };
}
