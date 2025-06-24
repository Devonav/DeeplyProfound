import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-head',
  templateUrl: './head.component.html',
  styleUrls: ['./head.component.css']
})
export class HeadComponent implements OnInit {
  showWelcome = true;
  isFadingOut = false;
  showMain = false;

  ngOnInit() {
    setTimeout(() => {
      this.isFadingOut = true;
      setTimeout(() => {
        this.showMain = true;
        this.showWelcome = false;
      }, 1500);
    }, 2000);
  }
}
