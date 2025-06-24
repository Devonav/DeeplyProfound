import { Component, OnInit } from '@angular/core';
import { TechService } from './tech.service';
import { Technology } from './tech.model';

@Component({
  selector: 'app-tech',
  templateUrl: './tech.component.html',
  styleUrls: ['./tech.component.css']
})
export class TechComponent implements OnInit {

  technologies: Technology[] = [];

  constructor(private readonly techService: TechService) { }

  ngOnInit() {
    this.techService.getTechnologies().subscribe((value: Technology[]) => {
      this.technologies = value;
    });
  }
}