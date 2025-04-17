import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideToastr } from 'ngx-toastr';
import { provideAnimations } from '@angular/platform-browser/animations';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideAnimations(),
    provideToastr({
      timeOut: 3000, // Toast will close after 3 seconds
      positionClass: 'toast-top-right', // Position of the toast
      preventDuplicates: true, // Prevent duplicate toasts
      progressBar: true, // Show progress bar
      closeButton: true, // Show close button
      tapToDismiss: false, // Disable tap to dismiss
    }), provideRouter(routes), provideHttpClient(withInterceptorsFromDi()),]
};
