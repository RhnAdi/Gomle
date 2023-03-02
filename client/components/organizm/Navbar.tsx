import dynamic from "next/dynamic";
import Logo from "../atoms/logo";
import ChatIcon from "../icons/Chat";
import HomeIcon from "../icons/Home";
import NotificationIcon from "../icons/Notification";
import SearchIcon from "../icons/Search";
import UserIcon from "../icons/user";
// import ThemeButton from "../moleculs/ThemeButton";
const ThemeButton = dynamic(() => import("../moleculs/ThemeButton"), {
  ssr: false,
});

export default function Navbar() {
  return (
    <>
      <div
        id="header"
        className="z-20 fixed w-full px-8 md:px-10 lg:px-16 py-3 bg-slate-50 dark:bg-slate-800 flex justify-between items-center shadow-sm gap-x-5"
      >
        <div id="logo">
          <Logo />
        </div>
        <div id="search" className="w-full md:w-96">
          {/* <Input placeholder="Search" icon={<SearchIcon />} /> */}
          <div className="relative flex items-center w-full">
            <input
              className="bg-slate-100 dark:bg-slate-700/50 w-full px-6 py-2 block rounded-xl outline-none focus:ring-2 focus:ring-sky-500 text-slate-700 dark:text-white focus:bg-slate-50 dark:focus:bg-slate-600"
              placeholder="Search"
            />
            <div className="text-slate-400 absolute right-6">
              <SearchIcon />
            </div>
          </div>
        </div>
        <div
          id="nav_info"
          className="justify-between items-center gap-x-8 flex"
        >
          <ThemeButton />
          <div id="home" className="text-sky-500 hidden md:block">
            <HomeIcon />
          </div>
          <div id="notification" className="text-slate-400 hidden md:block">
            <NotificationIcon />
          </div>
          <div id="chat" className="text-slate-400 hidden md:block">
            <ChatIcon />
          </div>
          <div id="account" className="text-slate-400 hidden md:block">
            <UserIcon />
          </div>
        </div>
      </div>
      <div className="z-20 px-16 py-4 flex justify-between items-center gap-x-8 fixed bottom-0 w-full md:hidden bg-slate-200">
        <div className="text-sky-500">
          <HomeIcon />
        </div>
        <div className="text-slate-500">
          <NotificationIcon />
        </div>
        <div className="text-slate-500">
          <ChatIcon />
        </div>
        <div className="text-slate-500">
          <UserIcon />
        </div>
      </div>
    </>
  );
}
