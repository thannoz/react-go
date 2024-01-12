import { ChangeEvent, useState } from "react";
import { useNavigate, useOutletContext } from "react-router-dom";
import Input from "./form/Input";

type OutletContext = {
  setJwtToken: (value: string) => void;
  setAlertMessage: (value: string) => void;
  setAlertClassName: (value: string) => void;
};

const Login = () => {
  const [email, setEmail] = useState<string>("");
  const [, setPassword] = useState<string>("");

  // Get props passed by the Outlet component in App.tsx
  const { setJwtToken, setAlertMessage, setAlertClassName } =
    useOutletContext() as OutletContext;
  const navigate = useNavigate();

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault(); // stop the form to submit

    if (email === "admin@example.com") {
      setJwtToken("abc");
      setAlertClassName("alert-success");
      setAlertMessage("Login was successful");
      // TODO: Delete function setAlertMessage in useEffect to make the message disappear
      navigate("/");
    } else {
      setAlertClassName("alert-danger");
      setAlertMessage("Invalid credentials");
    }
  };

  const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  return (
    <div className="col-md-6 offset-md-3">
      <h2>Login</h2>
      <hr />

      <form onSubmit={handleSubmit}>
        <Input
          title="Email Address"
          type="email"
          className="form-control"
          name="email"
          autoComplete="email-new"
          onChange={handleEmailChange}
          errorDiv={""}
          errorMsg={""}
        />
        <Input
          title="Password"
          type="password"
          className="form-control"
          name="password"
          autoComplete="password-new"
          onChange={handlePasswordChange}
          errorDiv={""}
          errorMsg={""}
        />
        <hr />
        <input className="btn btn-primary" type="submit" value="login" />
      </form>
    </div>
  );
};

export default Login;
