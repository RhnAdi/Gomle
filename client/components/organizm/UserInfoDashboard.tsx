import UserFollow from "../atoms/UserFollow";
import UserIcon from "../icons/user";
import Image from "next/image";
import useSWR from "swr";
import { getCookie } from "cookies-next";
import { getAccount } from "../../utils/fetcher";

export default function UserInfoDashboard() {
  const token = getCookie("auth");
  const { data, error } = useSWR(
    ["http://localhost:8080/api/v1/users/account/", token],
    getAccount
  );
  console.log(data);
  return (
    <div className="relative h-min pt-20">
      <div
        id="left"
        className="hidden md:flex md:right-10 lg:right-16 fixed overflow-scroll h-full md:w-[33%] lg:w-1/4 flex flex-col gap-y-4 top-0 pt-20"
      >
        <div className=" account_info bg-slate-50 dark:bg-slate-800 rounded-xl relative pb-4 shadow">
          <div className="z-0 banner relative w-full h-24 rounded-t-lg overflow-hidden">
            <Image
              src={
                data?.banner != ""
                  ? `http://127.0.0.1:8080/api/v1/public/images/${data?.banner}`
                  : "/auth_bg.jpg"
              }
              layout="fill"
              objectFit="cover"
              objectPosition="center"
              className="z-10"
            />
          </div>
          <div className="-mt-6 z-20 absolute w-full flex items-center justify-center">
            <div className="photo_profile p-2 rounded-full bg-slate-100 border border-sky-500 text-sky-500 w-min ring-2 ring-slate-50 dark:ring-slate-700 relative w-12 h-12 overflow-hidden flex items-center justify-center">
              {data?.photo_profile ? (
                <Image
                  alt="avatar_profile"
                  layout="fill"
                  src={`http://127.0.0.1:8080/public/images/data?.photo_profile`}
                />
              ) : (
                <UserIcon />
              )}
            </div>
          </div>
          <p className="text-slate-700 font-semibold font-body mt-8 text-center dark:text-white">
            {data?.username ? data.username : ""}
          </p>
          <div className="text-slate-500 dark:text-slate-300 flex gap-x-4 text-sm w-full items-center justify-center mt-2">
            <p>18 Post</p>
            <p>{data?.followers_count} Follower</p>
            <p>{data?.following_count} Following</p>
          </div>
        </div>
        <div className="followings_wrapper bg-slate-50 dark:bg-slate-800 px-4 py-5 rounded-xl flex gap-y-3 flex-col shadow">
          <p className="text-slate-700 dark:text-white text-md font-semibold">
            Followings
          </p>
          <div className="followings flex flex-col gap-y-3">
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
          </div>
        </div>
        <div className="followings_wrapper bg-slate-50 dark:bg-slate-800 px-4 py-5 rounded-xl flex gap-y-3 flex-col shadow">
          <p className="text-slate-700 dark:text-white text-md font-semibold">
            Followers
          </p>
          <div className="followings flex flex-col gap-y-3">
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
            <UserFollow />
          </div>
        </div>
      </div>
    </div>
  );
}
