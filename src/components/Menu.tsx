import MenuLink from "./MenuLink";

export default function Menu(): JSX.Element {
  return (
    <div className="flex justify-center items-center">
      <nav className="flex space-between">
        <MenuLink href="/" name="Home" />
      </nav>
    </div>
  )
}