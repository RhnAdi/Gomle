import Input from "../atoms/input";
import UserIcon from "../icons/user";
import HeartIcon from "../icons/Heart";
import ChatIcon from "../icons/Chat";
import { useEffect, useState } from "react";
import axios from "axios";
import * as React from "react";
import TimeAgo from "timeago-react";
import Gallery from "./Gallery";

type UserCardProps = {
  user_id: string;
  date: string;
  post: string;
  images?: string[];
};

type UserDetail = {
  UserID: string;
  username: string;
  firstname: string;
  lastname: string;
  photo_profile: string;
  banner: string;
};

export default function PostCard(props: UserCardProps) {
  const init: UserDetail = {
    UserID: "",
    username: "",
    firstname: "",
    lastname: "",
    photo_profile: "",
    banner: "",
  };
  const [userDetail, setUserDetail] = useState<UserDetail>(init);
  const [error, setError] = useState(false);
  useEffect(() => {
    async function fetch() {
      try {
        const res = await axios.get(`http://127.0.0.1:8080/users/${props.user_id}/detail`);
        const data = res.data.data;
        console.log(res);
        console.log(data);
        setUserDetail(data);
      } catch {
        setError(true);
      }
    }
    fetch();
  }, []);

  if (error) {
    return <p>Error</p>;
  }
  return (
    <div className="card bg-gray-50 dark:bg-gray-700 rounded-2xl overflow-hidden h-min w-full my-3 shadow p-1">
      <div className="user_info px-4 my-2 flex justify-between items-center gap-x-4">
        <div className="bg-gray-100 rounded-full p-2 text-sky-500 w-min border border-sky-500">
          <UserIcon />
        </div>
        <div className="flex-grow">
          <p className="text-gray-700 dark:text-white text-md font-semibold">{userDetail.username}</p>
          <p className="text-gray-500 dark:text-gray-300 text-sm font-medium">
            <TimeAgo datetime={props.date} locale="id_IDN" />
          </p>
        </div>
        <div className="flex gap-x-1 cursor-pointer p-2 hover:bg-gray-100 dark:hover:bg-gray-600 rounded-full">
          <div className="bg-gray-300 dark:bg-gray-400 rounded-full w-1.5 h-1.5"></div>
          <div className="bg-gray-300 dark:bg-gray-400 rounded-full w-1.5 h-1.5"></div>
          <div className="bg-gray-300 dark:bg-gray-400 rounded-full w-1.5 h-1.5"></div>
        </div>
      </div>
      {props.images && <Gallery files={props.images} />}
      <div className="flex gap-x-3 px-4 mt-4">
        <div className="flex gap-x-1">
          <HeartIcon className="text-pink-500" />
          <p className="text-gray-400">24</p>
        </div>
        <div className="flex gap-x-1">
          <ChatIcon className="text-gray-400" />
          <p className="text-gray-400">12</p>
        </div>
      </div>
      <div className="post px-4 mt-6 mb-4 text-gray-800 dark:text-gray-100 font-medium">
        <p>{props.post}</p>
      </div>
      <div className="px-4 border-t py-3 border-gray-100 dark:border-gray-600">
        <div className="flex gap-x-2 items-center">
          <div className="bg-gray-200 rounded-full p-2 text-sky-500 w-min">
            <UserIcon />
          </div>
          <Input placeholder="Comment" py={2} />
        </div>
      </div>
    </div>
  );
}
