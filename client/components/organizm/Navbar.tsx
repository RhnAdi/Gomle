import Logo from "../atoms/logo";
import ChatIcon from "../icons/Chat";
import HomeIcon from "../icons/Home";
import NotificationIcon from "../icons/Notification";
import SearchIcon from "../icons/Search";
import UserIcon from "../icons/user";
import ThemeButton from "../moleculs/ThemeButton";

export default function Navbar() {
  return (
    <>
      <div id="header" className="z-20 fixed w-full px-8 md:px-10 lg:px-16 py-3 bg-gray-50 dark:bg-gray-700 flex justify-between items-center shadow-sm gap-x-5">
        <div id="logo">
          <Logo />
        </div>
        <div id="search" className="w-full md:w-96">
          {/* <Input placeholder="Search" icon={<SearchIcon />} /> */}
          <div className="relative flex items-center w-full">
            <input
              className="bg-gray-100 dark:bg-gray-600 w-full px-6 py-2 block rounded-xl outline-none focus:ring-2 focus:ring-sky-500 text-gray-700 dark:text-white focus:bg-gray-50 dark:focus:bg-gray-600"
              placeholder="Search"
            />
            <div className="text-gray-400 absolute right-6">
              <SearchIcon />
            </div>
          </div>
        </div>
        <div id="nav_info" className="justify-between items-center gap-x-8 flex">
          <ThemeButton />
          <div id="home" className="text-sky-500 hidden md:block">
            <HomeIcon />
          </div>
          <div id="notification" className="text-gray-400 hidden md:block">
            <NotificationIcon />
          </div>
          <div id="chat" className="text-gray-400 hidden md:block">
            <ChatIcon />
          </div>
          <div id="account" className="text-gray-400 hidden md:block">
            <UserIcon />
          </div>
        </div>
      </div>
      <div className="z-20 px-16 py-4 flex justify-between items-center gap-x-8 fixed bottom-0 w-full md:hidden bg-gray-200">
        <div className="text-sky-500">
          <HomeIcon />
        </div>
        <div className="text-gray-400">
          <NotificationIcon />
        </div>
        <div className="text-gray-400">
          <ChatIcon />
        </div>
        <div className="text-gray-400">
          <UserIcon />
        </div>
      </div>
    </>
  );
}
