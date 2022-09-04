import Link from "next/link";
import auth_style from "../../styles/auth.module.css";
import Input from "../../components/atoms/input";
import AtIcon from "../../components/icons/At";
import PasswordInput from "../../components/atoms/PasswordInput";
import Logo from "../../components/atoms/logo";
import { useForm } from "react-hook-form";
import axios from "axios";
import { NextPageWithLayout } from "../_app";
import { ReactElement } from "react";
import { deleteCookie, hasCookie, setCookie } from "cookies-next";
import { useRouter } from "next/router";

const Login: NextPageWithLayout = () => {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    setError,
    formState: { isSubmitting, errors },
  } = useForm();
  const onSubmit = (data: any) => {
    return new Promise((resolve, reject) => {
      axios
        .post("http://127.0.0.1:8080/users/login", {
          email: data.email,
          password: data.password,
        })
        .then((res) => {
          const token = res.data.data.token;
          hasCookie("auth") && deleteCookie("auth");
          setCookie("auth", token);
          router.push("/");
          resolve(res);
        })
        .catch((err) => {
          setError("ErrorLogin", err.response.data.message);
          reject(err);
        });
    });
  };

  return (
    <div className={auth_style.wrapper + ` w-screen h-screen px-10 sm:px-16 md:px-20 lg:px-24 py-10`}>
      <div id="logo">
        <Logo />
      </div>
      <div id="content" className="text-white my-10">
        <p className="font-body font-bold">👋 Hello,</p>
        <p className="font-mono font-bold text-4xl">
          Welcome back<span className="text-blue-400">.</span>
        </p>
        <p className="text-body mt-5 font-light">
          Don't have an account ?{" "}
          <Link href="/auth/register">
            <span className="text-blue-400 cursor-pointer">Register here</span>
          </Link>
        </p>
      </div>
      <form onSubmit={handleSubmit(onSubmit)} id="register" className="font-body w-full md:w-96 flex flex-col gap-y-3">
        <Input config={register("email")} name="email" placeholder="Email" icon={<AtIcon />} type="email" autoComplete="off" />
        <PasswordInput config={register("password")} placeholder="Password" autoComplete="off" />
        <button
          type="submit"
          disabled={isSubmitting}
          className={`my-4 w-full md:w-96 bg-sky-600 text-white text-lg font-bold py-2 font-display rounded-3xl hover:bg-sky-700 ${
            isSubmitting ? "outline outline-offset-4 outline-2 outline-sky-700 shadow" : null
          }`}
        >
          {isSubmitting == true ? "Loading..." : "Login"}
        </button>
      </form>
    </div>
  );
};

Login.getLayout = (page: ReactElement) => {
  return <>{page}</>;
};

export default Login;
