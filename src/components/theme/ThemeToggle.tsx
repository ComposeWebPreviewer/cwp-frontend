import { useTheme } from "../../contexts/ThemeContext";
import './ThemeToggle.css'
import { Moon, Sun } from "lucide-react";

export default function ThemeToggle() {
  const { theme, toggleTheme } = useTheme();

  return (
    <button className="theme-toggle-button" onClick={toggleTheme}>
      {theme === 'light' ? (
        <Sun />
      ) : (
        <Moon />
      )}
    </button>
  );
}
