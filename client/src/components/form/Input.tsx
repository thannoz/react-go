import React, { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  title: string;
  errorDiv: string;
  errorMsg: string;
}

// eslint-disable-next-line react-refresh/only-export-components
const Input: React.ForwardRefRenderFunction<HTMLInputElement, InputProps> = (
  { title, errorDiv, errorMsg, ...rest },
  ref
) => {
  return (
    <div className="mb-3">
      <label className="form-label" htmlFor={rest.name}>
        {title}
      </label>
      <input {...rest} ref={ref} />
      <div className={errorDiv}>{errorMsg}</div>
    </div>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export default React.forwardRef(Input);
