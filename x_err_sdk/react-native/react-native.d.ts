// react-native 类型占位（SDK 运行时依赖宿主提供的 react-native，构建时 external 处理，
// 此处仅声明用到的 API，避免为编译引入重量级 react-native 类型包）
declare module "react-native" {
  export const Dimensions: {
    get(dim: "window" | "screen"): {
      width: number;
      height: number;
      scale: number;
      fontScale: number;
    };
  };
}
