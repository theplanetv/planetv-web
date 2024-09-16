import Link from "next/link";

type Props = {
	href: string
	name: string
}

export default function MenuLink(props: Props): JSX.Element {
  return (
    <Link className="px-5" href={props.href}>{props.name}</Link>
  )
}