import { useUserStore } from "@/stores/user";

export const perms = (perms) => {
  if (!perms) return true;
  const userStore = useUserStore();
  const auth = userStore.auth;

  let find = auth.find((item) => {
    return item.action === perms;
  });
  return !!find;
};
