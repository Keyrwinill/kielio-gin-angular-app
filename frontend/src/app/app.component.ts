import { Component, signal } from '@angular/core';
import { ApiService } from './services/api.service';

@Component({
  selector: 'app-root',
  imports: [],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class App {
  message = signal('Click the button to call Gin API');

  constructor(private apiService: ApiService) {}

  callApi() {
    this.apiService.getHealth().subscribe({
      next: (res) => this.message.set(res.message),
      error: () => this.message.set('Cannot connect to backend')
    });
  }
}
