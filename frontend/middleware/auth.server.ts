export default defineNuxtRouteMiddleware(async (to, from) => {
  type User = {
    email: string;
  };
  // query the server /refresh route to see if the user is logged in
  const { data, error } = await useFetch<User>(
    "http://localhost:8080/api/v1/refresh"
  );

  console.log(error.value);

  // if not, redirect to login
  if (error) {
    return navigateTo("/auth/login");
  }

  return;
});
