import { FC } from "react";
const ErrorPage: FC = () => {
  return (
    <div id="error-page">
      <h1>Oops!</h1>
      <p>Sorry, an unexpected error has occurred.</p>
      <p>
        <i>{}</i>
      </p>
    </div>
  );
};

export default ErrorPage;
