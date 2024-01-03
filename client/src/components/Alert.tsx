import { FC } from "react";

type AlterProps = {
  className: string;
  message: string;
};

export const Alert: FC<AlterProps> = ({ className, message }) => {
  return (
    <div className={`alert ${className}`} role="alert">
      {message}
    </div>
  );
};
