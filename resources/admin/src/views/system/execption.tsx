import { useRouteError } from "react-router-dom"
import { Button } from "antd"

export default function Component() {
    const error = useRouteError()

    console.dir(error)
    console.dir(error.status)
    return (
        <main className="flex size-full flex-col items-center justify-center rounded-lg p-4">
            <p className="m-3 text-sm font-semibold uppercase text-ink-faint">Error: 404</p>

            <h1 className="text-4xl font-bold">There's nothing here.</h1>

            <p className="mt-2 text-sm text-ink-dull">Its likely that this page has not been built yet, if so we're on it!</p>

            <Button>返回</Button>
        </main>
    )
}
