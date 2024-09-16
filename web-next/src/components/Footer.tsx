import Link from "next/link"
import GithubIcon from "./icons/GithubIcon"

export default function Footer(): JSX.Element {
  return (
    <footer className="flex justify-center items-center">
      <div className="flex-col">
        <p className="text-center">Follow me on</p>

        <div className="flex justify-center items-center">
          <Link
            href="https://github.com/goatastronaut0212"
            className="w-8 h-8"
          >
            <GithubIcon />
          </Link>
        </div>
      </div>
    </footer>
  )
}