import { Link, createLink } from '@tanstack/react-router'
import { Button, Dropdown, Menu, Navbar } from 'react-daisyui'
import { AiOutlineMenu } from 'react-icons/ai'

interface MenuItem {
  name: string
  href: string
}

interface MenuGroup {
  name: string
  items: MenuItem[]
}

type MenuData = (MenuItem | MenuGroup)[]

const menuData: MenuData = [
  {
    name: '主页',
    href: '/',
  },
  {
    name: '排行榜',
    href: '/rank',
  },
]

const DropdownItemRouter = createLink(Dropdown.Item)

const DropdownItems = ({ menuData }: { menuData: MenuData }) => {
  return (
    <>
      {menuData.map((item) => {
        if ('items' in item) {
          return (
            <Dropdown.Item key={item.name}>
              <details>
                <summary>{item.name}</summary>
                <ul className="p-2">
                  <DropdownItems menuData={item.items} />
                </ul>
              </details>
            </Dropdown.Item>
          )
        }
        return (
          <DropdownItemRouter key={item.name} to={item.href}>
            {item.name}
          </DropdownItemRouter>
        )
      })}
    </>
  )
}

const NavbarItems = ({ menuData }: { menuData: MenuData }) => {
  return (
    <>
      {menuData.map((item) => {
        if ('items' in item) {
          return (
            <Menu.Item key={item.name}>
              <details>
                <summary>{item.name}</summary>
                <ul className="p-2">
                  <NavbarItems menuData={item.items} />
                </ul>
              </details>
            </Menu.Item>
          )
        }
        return (
          <Menu.Item key={item.name}>
            <Link to={item.href}>{item.name}</Link>
          </Menu.Item>
        )
      })}
    </>
  )
}

export const AppBar = () => {
  return (
    <header className="sticky top-0 z-50 bg-base-100">
      <Navbar className="mx-auto max-w-6xl">
        <Navbar.Start>
          <Dropdown>
            <Button tag="label" color="ghost" tabIndex={0} className="md:hidden">
              <AiOutlineMenu />
            </Button>
            <Dropdown.Menu tabIndex={0} className="menu-sm z-[1] mt-3 w-52">
              <DropdownItems menuData={menuData} />
            </Dropdown.Menu>
          </Dropdown>
          <Link className="btn btn-ghost text-xl normal-case" to="/">
            GitPulse
          </Link>
        </Navbar.Start>
        <Navbar.End className="hidden md:flex">
          <Menu horizontal className="gap-2 px-1">
            <NavbarItems menuData={menuData} />
          </Menu>
        </Navbar.End>
      </Navbar>
    </header>
  )
}

export default AppBar
