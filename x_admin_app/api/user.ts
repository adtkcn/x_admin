import { request } from "@/utils/request";
// import type {Response} from '@/utils/request';

export function login(data: any) {
  return request<{
    token: string;
  }>({
    url: "/system/login",
    method: "POST",
    data,
  });
}

export function getInfo(token: string) {
  return request<{
    permissions: string[];
    user: {
      avatar: string;
      createTime: string;
      dept: string;
      id: number;
      isDisable: number;
      lastLoginIp: string;
      lastLoginTime: string;
      nickname: string;
      role: string;
      updateTime: string;
      username: string;
    };
  }>({
    url: "/system/admin/self",
    method: "GET",
    data: {
      token,
    },
  });
}

export function logout() {
  return request({
    url: "/system/logout",
    method: "GET",
  });
}
