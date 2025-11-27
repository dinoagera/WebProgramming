import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { App } from './app/app';
import { provideHttpClient } from '@angular/common/http'; // ← Angular 15+


bootstrapApplication(App, appConfig)
  .catch((err) => console.error(err));
