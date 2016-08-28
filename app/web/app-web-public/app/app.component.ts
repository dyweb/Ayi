import { Component } from '@angular/core';
// FIXME: vscode ts hint says path/to/hero.ts is not a module, but it is working in browser
import { Hero } from './hero';

const HEROES: Hero[] = [
  { id: 10, name: 'Tracer' },
  { id: 11, name: 'Winston' },
  { id: 12, name: 'Reaper' },
  { id: 13, name: 'D.V.A' },
  { id: 14, name: 'Hanzo' },
  { id: 15, name: 'Genji' },
  { id: 16, name: 'Mercy' },
  { id: 17, name: 'Pharah' },
  { id: 18, name: 'Road hog' },
  { id: 19, name: 'Anna' },
  { id: 20, name: 'Soilder 76' }
];

@Component({
  selector: 'my-app',
  template: `
  <h1>{{title}}</h1>
  <my-hero-detail [hero]="selectedHero"></my-hero-detail>
  <h2>My Heroes</h2>
  <ul class="heroes">
    <li *ngFor="let hero of heroes" 
        (click)="onSelect(hero)"
        [class.selected]="hero === selectedHero">
        <span class="badge">{{hero.id}}</span> {{hero.name}}
    </li>
  </ul>
  `,
  // FIXME: styles is not working -> I forgot to add class in template
  styles: [`
  .selected {
    background-color: #CFD8DC !important;
    color: white;
  }
  .heroes {
    margin: 0 0 2em 0;
    list-style-type: none;
    padding: 0;
    width: 15em;
  }
  .heroes li {
    cursor: pointer;
    position: relative;
    left: 0;
    background-color: #EEE;
    margin: .5em;
    padding: .3em 0;
    height: 1.6em;
    border-radius: 4px;
  }
  .heroes li.selected:hover {
    background-color: #BBD8DC !important;
    color: white;
  }
  .heroes li:hover {
    color: #607D8B;
    background-color: #DDD;
    left: .1em;
  }
  .heroes .text {
    position: relative;
    top: -3px;
  }
  .heroes .badge {
    display: inline-block;
    font-size: small;
    color: white;
    padding: 0.8em 0.7em 0 0.7em;
    background-color: #607D8B;
    line-height: 1em;
    position: relative;
    left: -1px;
    top: -4px;
    height: 1.8em;
    margin-right: .8em;
    border-radius: 4px 0 0 4px;
  }
`]
})
export class AppComponent {
    title = 'Overwatch';
    selectedHero: Hero;
    heroes = HEROES;
    onSelect(hero: Hero): void {
        this.selectedHero = hero;
    }
 }
