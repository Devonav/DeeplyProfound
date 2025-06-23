import { Component, AfterViewInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements AfterViewInit {
  title = 'go-trigger-app';

  ngAfterViewInit() {
    document.addEventListener('mousemove', (ev: MouseEvent) => {
      const cursor = document.getElementById('cursor');
      if (cursor) {
        cursor.style.top = `${ev.clientY}px`;
        cursor.style.left = `${ev.clientX}px`;
      }
    });
  }
}
