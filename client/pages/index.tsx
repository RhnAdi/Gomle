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
  created_at: string;
  updated_at: string;
};

export const getServerSideProps = async ({ req, res }: GetServerSidePropsContext) => {
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
};

const Home = ({ data }: InferGetServerSidePropsType<typeof getServerSideProps>) => {
  return (
    <>
      <div id="main_wrapper" className="relative flex flex-col md:flex-row jsutify-between gap-x-4">
        <div id="main" className="w-full md:w-[60%] lg:w-[70%] flex flex-col gap-y-0 md:gap-y-3 h-full ">
          <WritePost />
          <div className={`w-full items-center ${styles.post_wrapper}`}>
            {data.map((post, index) => {
              return <PostCard key={index} username={post.username} date={post.created_at} post={post.content} />;
            })}
          </div>
        </div>
        <UserInfoDashboard />
      </div>
    </>
  );
};

export default Home;
