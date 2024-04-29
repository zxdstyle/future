import { createEffect, createSignal, JSX, onMount } from 'solid-js';
import { Motion } from 'solid-motionone';
import { createScrollPosition } from '@solid-primitives/scroll';
import Lenis from 'lenis';

interface Props {
    children?: JSX.Element;
}

export default function SmoothScroll(props: Props) {
    let box: HTMLDivElement | undefined = undefined;
    onMount(() => {
        if (!box) return;
        const lenis = new Lenis({
            wrapper: box,
            lerp: 0.05,
            easing: x => 1 - Math.pow(1 - x, 4),
        });

        function raf(time) {
            lenis.raf(time);
            requestAnimationFrame(raf);
        }

        requestAnimationFrame(raf);
    });

    return (
        <div ref={box} class="h-full w-full overflow-auto">
            <div class=" h-full">{props.children}</div>
        </div>
    );
}
