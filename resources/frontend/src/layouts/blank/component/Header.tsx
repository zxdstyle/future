import { Motion } from 'solid-motionone';
import { useContext } from 'solid-js';
import { CursorContext } from '@/layouts/context/cursor';
import { A } from '@solidjs/router';

export default function Header() {
    const { setActive, setInactive } = useContext(CursorContext);

    return (
        <header class="fixed left-0 right-0 top-0 w-full z-top">
            <div class="flex justify-between p-12">
                <div class="logo" />
                <nav class="text-white-500 text-xs">
                    <ul class="flex gap-16">
                        <Motion.li
                            class="hover:text-white cursor-pointer"
                            onMouseOver={() => setActive()}
                            onMouseLeave={() => setInactive()}
                            initial={{ y: -10, x: -20, opacity: 0 }}
                            animate={{ y: 0, x: 0, opacity: 1 }}
                            transition={{ duration: 0.4, delay: 0 }}
                        >
                            <A href={'/'}>首页</A>
                        </Motion.li>
                        <Motion.li
                            class="hover:text-white cursor-pointer"
                            onMouseOver={() => setActive()}
                            onMouseLeave={() => setInactive()}
                            initial={{ y: -10, x: -20, opacity: 0 }}
                            animate={{ y: 0, x: 0, opacity: 1 }}
                            transition={{ duration: 0.4, delay: 0.2 }}
                        >
                            专辑
                        </Motion.li>
                    </ul>
                </nav>
            </div>
        </header>
    );
}
