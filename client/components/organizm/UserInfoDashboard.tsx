import UserFollow from "../atoms/UserFollow";
import UserIcon from "../icons/user";
import Image from "next/image";

export default function UserInfoDashboard() {
  return (
    <div id="left" className="hidden md:flex md:right-10 lg:right-16 fixed overflow-scroll h-full md:w-[33%] lg:w-1/4 flex flex-col gap-y-4">
      <div className="account_info bg-gray-50 dark:bg-gray-700 rounded-xl relative pb-4 shadow">
        <div className="z-0 banner relative w-full h-24 rounded-t-lg overflow-hidden">
          <Image src="/auth_bg.jpg" layout="fill" objectFit="cover" objectPosition="center" className="z-10" />
        </div>
        <div className="-mt-5 z-20 absolute w-full flex items-center justify-center">
          <div className="photo_profile p-2 rounded-full bg-gray-100 border border-sky-500 text-sky-500 w-min ring-2 ring-gray-50 dark:ring-gray-700">
            <UserIcon />
          </div>
        </div>
        <p className="text-gray-700 font-semibold font-body mt-6 text-center dark:text-white">Jhon Doe</p>
        <div className="text-gray-500 dark:text-gray-300 flex gap-x-4 text-sm w-full items-center justify-center">
          <p>18 Post</p>
          <p>40 Follower</p>
          <p>65 Following</p>
        </div>
      </div>
      <div className="followings_wrapper bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-xl flex gap-y-3 flex-col shadow">
        <p className="text-gray-700 dark:text-white text-md font-semibold">Followings</p>
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
      <div className="followings_wrapper bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-xl flex gap-y-3 flex-col shadow">
        <p className="text-gray-700 dark:text-white text-md font-semibold">Followers</p>
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
  );
}
