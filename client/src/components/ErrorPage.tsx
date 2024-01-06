import { FC } from "react";
import { useRouteError } from "react-router-dom";

type CustomRouteError = Error | { status: number; message: string };

const ErrorPage: FC = () => {
  const err = useRouteError() as CustomRouteError;

  if (err instanceof Error) {
    return <div>Error: {err.message}</div>;
  } else {
    return (
      <div>
        <h1>Oops!</h1>
        <h1>Sorry, an unexpected error has occurred.</h1>
        <h1>Status: {err.status}</h1>
        <h2>Page not found</h2>
      </div>
    );
  }
};

export default ErrorPage;
