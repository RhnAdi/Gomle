import MoonIcon from "../icons/Moon";
import { useTheme } from "next-themes";
import SunIcon from "../icons/Sun";

export default function ThemeButton() {
  const { theme, setTheme } = useTheme();
  return (
    <>
      <button title="theme_button" className="text-yellow-500" onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
        {theme == "dark" ? <SunIcon /> : <MoonIcon />}
      </button>
    </>
  );
}
