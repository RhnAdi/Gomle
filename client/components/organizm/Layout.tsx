import Navbar from "./Navbar";

export default function Layout(props: any) {
  return (
    <div className="overflow-hidden">
      <Navbar />
      <div id="wrapper" className="w-screen pt-20 px-2 md:px-10 lg:px-16">
        {props.children}
      </div>
    </div>
  );
}
