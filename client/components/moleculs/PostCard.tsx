import Image from "next/image";
import Input from "../atoms/input";
import UserIcon from "../icons/user";
import HeartIcon from "../icons/Heart";
import ChatIcon from "../icons/Chat";

type UserCardProps = {
  username: string;
  date: string;
  post: string;
  image?: string;
};

export default function PostCard(props: UserCardProps) {
  return (
    <div className="card bg-gray-50 dark:bg-gray-700 rounded-2xl overflow-hidden h-min w-full my-3 shadow p-1">
      <div className="user_info px-4 my-2 flex justify-between items-center gap-x-4">
        <div className="bg-gray-100 rounded-full p-2 text-sky-500 w-min border border-sky-500">
          <UserIcon />
        </div>
        <div className="flex-grow">
          <p className="text-gray-700 dark:text-white text-md font-semibold">{props.username}</p>
          <p className="text-gray-500 dark:text-gray-300 text-sm font-medium">{props.date}</p>
        </div>
        <div className="flex gap-x-1 cursor-pointer p-2 hover:bg-gray-100 dark:hover:bg-gray-600 rounded-full">
          <div className="bg-gray-300 dark:bg-gray-400 rounded-full w-1.5 h-1.5"></div>
          <div className="bg-gray-300 dark:bg-gray-400 rounded-full w-1.5 h-1.5"></div>
          <div className="bg-gray-300 dark:bg-gray-400 rounded-full w-1.5 h-1.5"></div>
        </div>
      </div>
      {props.image ? (
        <div className="image mt-2 mb-6 relative w-full h-64 rounded-lg overflow-hidden">
          <Image src={props.image} layout="fill" />
        </div>
      ) : null}
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
