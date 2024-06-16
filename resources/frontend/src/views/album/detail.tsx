import NavLine from '@/components/NavLine';
import { CursorContext } from '@/layouts/context/cursor';
import { useContext } from 'solid-js';
import Carousel from '@/views/album/components/Carousel';
import Masonry from '@/views/album/components/Masonry';

export default function Detail() {
    // const images = [
    //     'https://images.pexels.com/photos/6954544/pexels-photo-6954544.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/6954509/pexels-photo-6954509.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/11153461/pexels-photo-11153461.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/5049271/pexels-photo-5049271.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    //     'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    // ];

    const { setActive, setInactive } = useContext(CursorContext);

    return (
        <div class="w-full h-full flex bg-black/90 z-top">
            <NavLine title="焦糖拿铁" subTitle={'Caramel latte'} class="left-36" position="center" />

            <Masonry />

            {/*<Carousel />*/}

            <NavLine title="返回" subTitle="return" class="right-36" position="bottom" onClick={() => history.back()} />
        </div>
    );
}
