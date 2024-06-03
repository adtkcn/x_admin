import { defineStore } from "pinia";
import { logout, getInfo } from "@/api/user";

import { getLocalStorage, removeLocalStorage } from "@/utils/storage.js";

export const useUserStore = defineStore("user", {
  state: () => {
    return {
      nickname: "",
      username: "",
      avatar: "",
      userInfo: {},
      menu: [],
      auth: [],
    };
  },
  // 也可以这样定义
  // state: () => ({ count: 0 })
  actions: {
    // get user info
    getInfo() {
      return new Promise((resolve, reject) => {
        let token = getLocalStorage("token");
        if (!token) {
          return reject(new Error("需要登录"));
        }
        getInfo(token)
          .then((res) => {
            if (res.code === 200) {
              const { data } = res;
              this.nickname = data.user.nickname;
              this.username = data.user.username;
              
              this.avatar = data.user.avatar ||"";
              // this.menu = data.menu;
              // this.auth = data.auth;
              this.userInfo = data.user;
              resolve(data);
            } else {
              reject(new Error(res.message));
            }

            //   addRoutes(result.menu);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },

    // user logout
    logout() {
      return new Promise((resolve, reject) => {
        logout()
          .then(() => {
            //   resetRouter();
            this.resetToken();
            resolve();
          })
          .catch((error) => {
            reject(error);
          });
      });
    },

    // remove token
    resetToken() {
      return new Promise((resolve) => {
        // this.token = "";
        this.name = "";
        this.avatar = "";
        removeLocalStorage("token");
        resolve();
      });
    },
  },
});
