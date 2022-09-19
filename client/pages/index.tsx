import type { NextPage, GetServerSidePropsContext, GetServerSideProps, InferGetServerSidePropsType } from "next";
import PostCard from "../components/moleculs/PostCard";
import WritePost from "../components/moleculs/WritePost";
import styles from "../styles/Home.module.css";
import UserInfoDashboard from "../components/organizm/UserInfoDashboard";
import axios from "axios";
import { getCookie } from "cookies-next";

type Post = {
  id: string;
  user_id: string;
  username: string;
  content: string;
  files: string[];
  created_at: string;
  updated_at: string;
};

export const getServerSideProps = async ({ req, res }: GetServerSidePropsContext) => {
  try {
    const token = getCookie("auth", { req, res });
    const response = await axios.get("http://127.0.0.1:8080/post/dashboard", {
      headers: {
        Authorization: `${token}`,
      },
    });
    const data: Post[] = response.data.data;

    return {
      props: {
        data,
      },
    };
  } catch (err) {
    return {
      props: {
        data: null,
      },
      redirect: {
        destination: "/auth/login",
      },
    };
  }
};

const Home = ({ data }: InferGetServerSidePropsType<typeof getServerSideProps>) => {
  console.log(data);
  return (
    <>
      <div id="main_wrapper" className="relative flex flex-col md:flex-row jsutify-between gap-x-4">
        <div id="main" className="w-full md:w-[60%] lg:w-[70%] flex flex-col gap-y-0 md:gap-y-3 h-full ">
          <WritePost />
          <div className={`w-full items-center ${styles.post_wrapper}`}>
            {data?.map((post, index) => {
              return <PostCard key={index} date={post.created_at} post={post.content} user_id={post.user_id} images={post.files} />;
            })}
          </div>
        </div>
        <UserInfoDashboard />
      </div>
    </>
  );
};

export default Home;
