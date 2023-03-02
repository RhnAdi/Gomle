import { Swiper, SwiperSlide } from "swiper/react";

// Import Swiper styles
import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";

// import required modules
import { Navigation, Pagination, Mousewheel, Keyboard } from "swiper";
import ArrowRightIcon from "../icons/ArrowRight";
import ArrowLeftIcon from "../icons/ArrowLeft";
import { useRef } from "react";

type GalleryProps = {
  files: any[];
};

export default function Gallery(props: GalleryProps) {
  const { files } = props;

  return (
    <>
      <Swiper
        cssMode={true}
        navigation={{
          prevEl: ".prev",
          nextEl: ".next",
        }}
        pagination={true}
        mousewheel={true}
        keyboard={true}
        modules={[Navigation, Pagination, Mousewheel, Keyboard]}
        className="bg-slate-100 dark:bg-slate-700/50 relative"
      >
        {files.map((file, idx) => {
          return (
            <SwiperSlide tabIndex={idx} className="my-auto" key={idx}>
              <img
                alt={`file-${file.filename}`}
                src={`http://127.0.0.1:8080/public/images/${file.filename}`}
                className="w-full h-auto"
              />
            </SwiperSlide>
          );
        })}
        <ArrowRightIcon
          width={28}
          height={28}
          className="prev absolute top-1/2 left-2 text-slate-800 bg-slate-200 dark:bg-white rounded-full z-50 hover:bg-slate-300"
        />
        <div className="next absolute top-1/2 right-2 text-slate-800 bg-slate-200 dark:bg-white rounded-full hover:bg-slate-300 z-50 cursor-pointer">
          <ArrowLeftIcon width={28} height={28} />
        </div>
      </Swiper>
    </>
  );
}
