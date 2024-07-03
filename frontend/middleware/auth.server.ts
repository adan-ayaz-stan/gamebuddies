export default defineNuxtRouteMiddleware(async (to, from) => {
  type User = {
    email: string;
  };
  // query the server /refresh route to see if the user is logged in
  let cookie = useCookie("gamebuddy");
  const { data, error } = await useFetch<User>(
    "http://localhost:8080/api/v1/refresh",
    {
      credentials: "include",
      headers: {
        Cookie: cookie.value != null ? `gamebuddy=${cookie.value}` : "",
      },
    }
  );

  // if not, redirect to login
  if (error.value != null) {
    return navigateTo("/auth/login");
  }

  const authPaths = ["/auth/login", "auth/register"];

  // if user is not logged in and not on the login page or the register page
  if (error.value && !authPaths.includes(to.path)) {
    return navigateTo("/auth/login");
  }

  // if user is logged in and on the login page or the register page
  if (!error.value && authPaths.includes(to.path)) {
    return navigateTo("/dashboard");
  }
});
