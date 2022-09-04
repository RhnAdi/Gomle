import Input from "../atoms/input";
import PaperPlaneIcon from "../icons/PaperPlane";
import UserIcon from "../icons/user";

export default function WritePost() {
  return (
    <div className="bg-gray-50 dark:bg-gray-700 w-full p-4 rounded-2xl shadow-sm">
      <div className="flex gap-x-3 items-center">
        <div className="bg-gray-100 dark:bg-gray-500 rounded-full p-2 border border-sky-500">
          <UserIcon className="text-sky-500" />
        </div>
        <Input placeholder="Write content ..." />
        <button type="button" title="posting_btn" className="bg-sky-500 py-3 px-3 rounded-full text-white flex gap-x-2">
          <PaperPlaneIcon />
        </button>
      </div>
    </div>
  );
}
