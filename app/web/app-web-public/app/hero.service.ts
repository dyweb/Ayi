import { Injectable } from '@angular/core';
import { HEROES } from './mock-heros';
import { Hero } from './hero';

@Injectable()
export class HeroService {
    getHeroes(): Promise<Hero[]> {
        console.log('service called!');
        console.log（HEROES);
        return Promise.resolve(HEROES);
    }
}