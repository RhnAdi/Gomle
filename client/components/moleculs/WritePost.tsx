import Input from "../atoms/input";
import ImageIcon from "../icons/Image";
import PaperPlaneIcon from "../icons/PaperPlane";
import UserIcon from "../icons/user";

export default function WritePost() {
  return (
    <div className="bg-gray-50 dark:bg-gray-700 w-full p-4 rounded-2xl shadow-sm">
      <div className="flex gap-x-3 items-center">
        <div className="bg-gray-100 dark:bg-gray-500 rounded-full p-2 border border-sky-500">
          <UserIcon className="text-sky-500" />
        </div>
        {/* <Input placeholder="Write content ..." /> */}
        <div className="relative flex items-center w-full gap-x-3">
          <textarea
            className={`py-3 bg-gray-100 text-gray-800 dark:bg-gray-600 w-full px-6 block rounded-xl outline-none focus:ring-2 focus:ring-sky-500 dark:text-white focus:bg-gray-50 dark:focus:bg-gray-600 h-12`}
            placeholder="Write Post..."
          />
          <label htmlFor="files" className="flex items-center">
            <div className="text-gray-500 dark:text-gray-300">
              <ImageIcon width={28} height={28} />
            </div>
          </label>
          <input id="files" type="file" placeholder="Input Images" className="hidden" />
        </div>
        <button type="button" title="posting_btn" className="bg-sky-500 py-3 px-3 rounded-full text-white">
          <PaperPlaneIcon />
        </button>
      </div>
    </div>
  );
}
