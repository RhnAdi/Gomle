import Link from "next/link";
import auth_style from "../../styles/auth.module.css";
import UserIcon from "../../components/icons/user";
import InfoIcon from "../../components/icons/info";
import Input from "../../components/atoms/input";
import AtIcon from "../../components/icons/At";
import PasswordInput from "../../components/atoms/PasswordInput";
import Logo from "../../components/atoms/logo";
import { useForm } from "react-hook-form";
import axios from "axios";
import { ReactElement } from "react";
import { NextPageWithLayout } from "../_app";
import { deleteCookie, hasCookie, setCookie } from "cookies-next";
import { useRouter } from "next/router";

const Register: NextPageWithLayout = () => {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { isSubmitting },
  } = useForm();
  const onSubmit = (data: any) => {
    return new Promise((resolve, reject) => {
      axios
        .post("http://127.0.0.1:8080/users/register", {
          firstname: data.firstname,
          lastname: data.lastname,
          username: data.username,
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
          console.log(err);
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
        <p className="font-body font-bold">Let's go,</p>
        <p className="font-mono font-bold text-4xl">
          Create new account<span className="text-blue-400">.</span>
        </p>
        <p className="text-body mt-5 font-light">
          Already have account?{" "}
          <Link href="/auth/login">
            <span className="text-blue-400 cursor-pointer">Login here</span>
          </Link>
        </p>
      </div>
      <form id="register" onSubmit={handleSubmit(onSubmit)} className="font-body w-full md:w-96 flex flex-col gap-y-3">
        <div className="flex justify-between gap-x-3 w-full">
          <Input config={register("firstname")} name="firstname" placeholder="Firstname" icon={<InfoIcon />} />
          <Input config={register("lastname")} name="lastname" placeholder="Lastname" icon={<InfoIcon />} />
        </div>
        <Input config={register("username")} name="username" placeholder="Username" icon={<UserIcon />} autoComplete="off" />
        <Input config={register("email")} name="email" placeholder="Email" icon={<AtIcon />} type="email" autoComplete="off" />
        <PasswordInput config={register("password")} name="password" placeholder="Password" autoComplete="off" />
        <button
          type="submit"
          disabled={isSubmitting}
          className={`my-4 w-full md:w-96 bg-sky-600 text-white text-lg font-bold py-2 font-display rounded-3xl hover:bg-sky-700 ${
            isSubmitting ? "outline outline-offset-4 outline-2 outline-sky-700 shadow" : null
          }`}
        >
          {isSubmitting == true ? "Loading..." : "Register"}
        </button>
      </form>
    </div>
  );
};

Register.getLayout = (page: ReactElement) => {
  return <>{page}</>;
};

export default Register;
