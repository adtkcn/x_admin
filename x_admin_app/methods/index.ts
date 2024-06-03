import { perms } from "@/utils/perms.js";
import { toPath } from "@/utils/utils.js";
import env from "@/utils/env";

/**
 * @description: 获取文件路径
 * @param {String} src 文件路径
 * @return {String} 文件路径
 */
export function filePath(src :string) { 
  let path = "";
  if (src) {
    if (src.startsWith("http")) {
      return src;
    }
    path = `${env.baseUrl}${src}`;
  }
  return path;
}


// 在全局配置对象中添加$perms属性，并将其值设为perms
export function installMethods(app) {
  app.config.globalProperties.$perms = perms;
  app.config.globalProperties.$toPath = toPath;
  app.config.globalProperties.$filePath = filePath;
}
