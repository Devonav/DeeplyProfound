import { Component, Input, HostListener } from '@angular/core';

@Component({
  selector: 'app-project-card',
  templateUrl: './project-card.component.html',
  styleUrls: ['./project-card.component.css']
})
export class ProjectCardComponent {
  @Input() name!: string;
  @Input() github?: string;
  @Input() website?: string;
  @Input() npm?: string;
  @Input() image!: string;

  @HostListener('mousemove', ['$event'])
  handleCardBackground(event: MouseEvent) {
    // Get the card element itself
    const target = event.target as HTMLElement;
    const rect = target.getBoundingClientRect();
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;

    target.style.setProperty('--y', `${y}px`);
    target.style.setProperty('--x', `${x}px`);
  }
}
