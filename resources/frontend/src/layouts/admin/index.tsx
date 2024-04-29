import { RouteSectionProps } from '@solidjs/router';

export default function Admin(props: RouteSectionProps) {
    return (
        <div class="flex">
            <aside class="w-72">aside</aside>
            <div>
                <header>header</header>

                <main>{props.children}</main>

                <footer>footer</footer>
            </div>
        </div>
    );
}
