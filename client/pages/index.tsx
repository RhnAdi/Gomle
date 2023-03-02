import type {
  GetServerSidePropsContext,
  InferGetServerSidePropsType,
} from "next";
const PostCard = dynamic(() => import("../components/moleculs/PostCard"), {
  ssr: false,
});
import WritePost from "../components/moleculs/WritePost";
import styles from "../styles/Home.module.css";
import UserInfoDashboard from "../components/organizm/UserInfoDashboard";
import axios from "axios";
import { getCookie } from "cookies-next";
import dynamic from "next/dynamic";
import useSWR from "swr";

type Post = {
  id: string;
  user_id: string;
  username: string;
  content: string;
  files: string[];
  created_at: string;
  updated_at: string;
};

export const getServerSideProps = async ({
  req,
  res,
}: GetServerSidePropsContext) => {
  try {
    const token = getCookie("auth", { req, res });
    const response = await axios.get(
      "http://127.0.0.1:8080/api/v1/post/dashboard",
      {
        headers: {
          Authorization: `${token}`,
        },
      }
    );
    const data: Post[] = response.data.data;

    return {
      props: {
        init_data: data,
      },
    };
  } catch (err) {
    return {
      props: {
        init_data: null,
      },
      redirect: {
        destination: "/auth/login",
      },
    };
  }
};

const Home = ({
  init_data,
}: InferGetServerSidePropsType<typeof getServerSideProps>) => {
  const token = getCookie("auth");
  const fetcher = (url: string) =>
    axios
      .get(url, {
        headers: {
          Authorization: `${token}`,
        },
      })
      .then((res) => res.data.data);
  const { data, error } = useSWR(
    "http://127.0.0.1:8080/api/v1/post/dashboard",
    fetcher,
    init_data ? { fallbackData: init_data } : {}
  );

  return (
    <>
      {error ? <p>Error</p> : null}
      <div
        id="main_wrapper"
        className="relative h-full flex flex-col md:flex-row jsutify-between gap-x-4"
      >
        <div
          id="main"
          className="w-full md:w-[60%] lg:w-[70%] flex flex-col gap-y-0 md:gap-y-3 h-full "
        >
          <WritePost />
          <div className={`w-full items-center ${styles.post_wrapper}`}>
            {data?.map((post: Post, index: number) => {
              return (
                <PostCard
                  key={index}
                  date={post.created_at}
                  post={post.content}
                  user_id={post.user_id}
                  images={post.files}
                />
              );
            })}
          </div>
        </div>
        <UserInfoDashboard />
      </div>
    </>
  );
};

export default Home;
