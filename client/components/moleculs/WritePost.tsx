import { rejects } from "assert";
import axios from "axios";
import { getCookie } from "cookies-next";
import { useForm } from "react-hook-form";
import Textarea from "../atoms/textarea";
import ImageIcon from "../icons/Image";
import PaperPlaneIcon from "../icons/PaperPlane";
import UserIcon from "../icons/user";

export default function WritePost() {
  const { register, handleSubmit } = useForm();
  const onSubmit = (data: any) => {
    const formData = new FormData();
    formData.set("content", data.content);
    Array.from(data.file).forEach((f: any) => {
      formData.append("files", f);
    });
    return new Promise((resolve, reject) => {
      const token = getCookie("auth");
      console.log(data);
      axios
        .post("http://127.0.0.1:8080/api/v1/post/", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
            Authorization: `${token}`,
          },
        })
        .then((res) => {
          console.log(res);
          resolve(res);
        })
        .catch((err) => {
          console.log(err);
          reject(err);
        });
    });
  };
  return (
    <div className="bg-slate-50 dark:bg-slate-800 w-full px-3 py-2 rounded-full shadow-sm">
      <form onSubmit={handleSubmit(onSubmit)} encType="multipart/form-data">
        <div className="flex gap-x-3 items-center">
          <div className="bg-slate-100 dark:bg-slate-500 rounded-full p-2 border border-sky-500">
            <UserIcon className="text-sky-500" />
          </div>
          <div className="relative flex items-center w-full gap-x-3">
            <Textarea
              config={register("content")}
              placeholder="Write Post..."
            />
            <label htmlFor="files" className="flex items-center">
              <div className="text-slate-500 dark:text-slate-300">
                <ImageIcon width={28} height={28} />
              </div>
            </label>
            <input
              {...register("file")}
              id="files"
              type="file"
              multiple={true}
              placeholder="Input Images"
              className="hidden"
            />
          </div>
          <button
            type="submit"
            title="posting_btn"
            className="bg-sky-500 py-3 px-3 rounded-full text-white"
          >
            <PaperPlaneIcon />
          </button>
        </div>
      </form>
    </div>
  );
}
