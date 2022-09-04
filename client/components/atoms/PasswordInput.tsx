import { ReactElement, HTMLProps, useState } from "react";
import { UseFormRegisterReturn } from "react-hook-form";
import EyeIcon from "../icons/eye";
import EyeOffIcon from "../icons/EyeOff";

interface PasswordInputProps extends HTMLProps<HTMLInputElement> {
  icon?: ReactElement;
  config: UseFormRegisterReturn;
  py?: string | number;
}

export default function PasswordInput(props: PasswordInputProps) {
  const [show, setShow] = useState(false);

  const handleShow = () => {
    setShow(!show);
  };
  return (
    <div className="relative flex items-center w-full">
      <input
        id={props.name}
        placeholder={props.name}
        className={`${
          props.py ? "py-" + props.py.toString() : "py-3"
        } bg-gray-100 text-gray-800 dark:bg-gray-600 w-full px-6 block rounded-xl outline-none focus:ring-2 focus:ring-sky-500 dark:text-white focus:bg-gray-50 dark:focus:bg-gray-600`}
        type={show ? "string" : "password"}
        {...props.config}
        {...props}
      />
      <div className="text-gray-400 absolute right-6">{show ? <EyeOffIcon onClick={handleShow} /> : <EyeIcon onClick={handleShow} />}</div>
    </div>
  );
}
