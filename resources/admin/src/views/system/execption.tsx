import type { AbortedDeferredError } from "react-router-dom"
import { useNavigate, useRouteError } from "react-router-dom"
import { Button } from "antd"

export default function Component() {
    const error = useRouteError() as AbortedDeferredError & { status?: number }
    const navigate = useNavigate()

    return (
        <main className="flex size-full flex-col items-center justify-center rounded-lg p-4 text-center">
            <p className="m-3 text-sm font-semibold uppercase text-ink-faint">
                Error:
                {error?.status ?? "404"}
            </p>

            <h1 className="text-4xl font-bold text-white">{error?.message ?? "There's nothing here."}</h1>

            <p className="mt-2 text-sm text-ink-dull">{error?.stack}</p>

            <Button onClick={() => navigate(-1)}>返回</Button>
        </main>
    )
}
