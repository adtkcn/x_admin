import { useUserStore } from "@/stores/user";

export const perms = (perms:string) => {
  if (!perms) return false;
  const userStore = useUserStore();
  const permissions = userStore.permissions;

  let find = permissions.find((item) => {
    return item === perms|| item === "*";
  });
  return !!find;
};
