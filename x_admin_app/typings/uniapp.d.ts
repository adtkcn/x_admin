/// <reference types="@dcloudio/types" />
/// <reference types="@dcloudio/types" />
/// <reference types="@dcloudio/types" />
/// <reference types="@dcloudio/types" />

declare module '@dcloudio/uni-app' {
export  const onShow: (callback: ((options?: App.LaunchShowOption) => void) | (() => void), target?:  null) => Function;
export  const onHide: (callback: (() => void) | (() => void), target?:  null) => Function;
export  const onLaunch: (callback: (options?: App.LaunchShowOption) => void, target?:  null) => Function;
export  const onError: (callback: (error: string) => void, target?:  null) => Function;
export  const onPageNotFound: (callback: (options: App.PageNotFoundOption) => void, target?:  null) => Function;
export  const onUnhandledRejection: (callback: (options: UniApp.OnUnhandledRejectionCallbackResult) => void, target?:  null) => Function;
export  const onThemeChange: (callback: (options: UniApp.OnThemeChangeCallbackResult) => void, target?:  null) => Function;
export  const onUniNViewMessage: (callback: (options: AnyObject) => void, target?:  null) => Function;
export  const onInit: (callback: (query?: AnyObject) => void, target?:  null) => Function;
export  const onLoad: (callback: (query?: AnyObject) => void, target?:  null) => Function;
export  const onReady: (callback: () => void, target?:  null) => Function;
export  const onUnload: (callback: () => void, target?:  null) => Function;
export  const onPullDownRefresh: (callback: () => void, target?:  null) => Function;
export  const onReachBottom: (callback: () => void, target?:  null) => Function;
export  const onShareAppMessage: (callback: (options: Page.ShareAppMessageOption) => Page.CustomShareContent, target?:  null) => Function;
export  const onShareTimeline: (callback: () => Page.ShareTimelineContent, target?:  null) => Function;
export  const onAddToFavorites: (callback: (options: Page.AddToFavoritesOption) => Page.CustomFavoritesContent, target?:  null) => Function;
export  const onPageScroll: (callback: (options: Page.PageScrollOption) => void, target?:  null) => Function;
export  const onResize: (callback: (options: Page.PageScrollOption) => void, target?:  null) => Function;
export  const onTabItemTap: (callback: (options: Page.TabItemTapOption) => void, target?:  null) => Function;
export  const onNavigationBarButtonTap: (callback: (options: Page.NavigationBarButtonTapOption) => void, target?:  null) => Function;
export  const onBackPress: (callback: (options: Page.BackPressOption) => any, target?:  null) => Function;
export  const onNavigationBarSearchInputChanged: (callback: (event: Page.NavigationBarSearchInputEvent) => void, target?:  null) => Function;
export  const onNavigationBarSearchInputConfirmed: (callback: (event: Page.NavigationBarSearchInputEvent) => void, target?:  null) => Function;
export  const onNavigationBarSearchInputClicked: (callback: () => void, target?:  null) => Function;
}