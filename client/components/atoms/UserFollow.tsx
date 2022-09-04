import UserIcon from "../icons/user";

export default function UserFollow() {
  return (
    <div className="flex gap-x-3 items-center">
      <div className="bg-gray-100 rounded-full p-1.5 text-sky-500 w-min border border-sky-500">
        <UserIcon />
      </div>
      <div className="flex-grow">
        <p className="text-gray-700 dark:text-white text-sm font-semibold">Angelina Catarina</p>
        <p className="text-gray-500 dark:text-gray-300 text-sm font-medium">Web Developer</p>
      </div>
      <div className="bg-sky-200/70 px-3 py-1 rounded-full text-sm dark:bg-sky-600/20">
        <p className="text-sky-600 dark:text-sky-300">Follow</p>
      </div>
    </div>
  );
}
