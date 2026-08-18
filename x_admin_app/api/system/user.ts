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
      create_time: string;
      dept: string;
      id: string;
      is_disable: number;
      last_login_ip: string;
      last_login_time: string;
      nickname: string;
      role: string;
      update_time: string;
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
