function HomePage() {
  return <div>{`Welcome to ${process.env.NEXT_PUBLIC_SOME_ENV_VAR}`}</div>
}

export default HomePage