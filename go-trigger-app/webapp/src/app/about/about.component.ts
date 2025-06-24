import { AfterViewInit, Component } from '@angular/core';
// @ts-ignore
import cssParticles from 'css-particles';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent implements AfterViewInit {
  // count = 0; // for like button if needed

  ngAfterViewInit() {
    this.moveImage();

    const card = document.getElementById('about-card') as HTMLElement;
    const canvas = document.getElementById('canvas') as HTMLCanvasElement;

    if (!canvas || !card) return;

    canvas.height = card.clientHeight;
    canvas.width = card.clientWidth;

    // Initialize css-particles if available
    // @ts-ignore
    const system = new cssParticles.ParticleSystem(canvas, { x: canvas.width, y: canvas.height });
    system.ammount = 100;
    system.speed = { x: { min: -5, max: 5 }, y: { min: -5, max: 5 } };
    system.diameter = { min: 1, max: 4 };
    system.life = { min: 5, max: 15 };
    system.init();
  }

  moveImage() {
    const height = document.body.clientHeight / 2;
    const group = document.getElementById('yo') as HTMLElement;
    if (!group) return;
    document.addEventListener('mousemove', (e: MouseEvent) => {
      const degrees = (height - e.clientY) * 5 / height;
      group.style.transform = `rotate(${degrees}deg)`;
    });
  }
}
