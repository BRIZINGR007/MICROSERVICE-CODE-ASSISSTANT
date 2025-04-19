import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../../core/services/dashboard.service';
import { CodebasePayload } from '../../core/interfaces/dashboard.interfaces';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgClass } from '@angular/common';
import { ToastrService } from 'ngx-toastr';
import { HeaderComponent } from '../header/header.component';
import { ChatService } from '../../core/services/chat.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  imports: [NgClass, ReactiveFormsModule, HeaderComponent],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.scss'
})
export class DashboardComponent implements OnInit {
  codebases: CodebasePayload[] = [];
  syncYourCodeBaseClicked: boolean = false;
  codebaseForm: FormGroup;
  submitted = false;
  constructor(
    private chatService: ChatService,
    private dashBoardService: DashboardService,
    private fb: FormBuilder,
    private toastr: ToastrService,
    private router: Router,
  ) {
    this.codebaseForm = this.fb.group({
      codeBaseName: ['', Validators.required],
      gitHubURL: ['', Validators.required],
      username: ['', Validators.required],
      token: ['', Validators.required],
      branch: ['', Validators.required],
      folderPath: ['', Validators.required],
    });
  }
  ngOnInit() {
    this.getCodeBases()
  }
  get f() {
    return this.codebaseForm.controls;
  }

  private async getCodeBases() {
    this.codebases = await this.dashBoardService.retrieveCodeBaseData();
    console.log(`CodeBase Data : ${this.codebases}`)
  }
  syncYourCodeBase() {
    this.syncYourCodeBaseClicked = true;
  }
  onSubmit() {
    this.submitted = true;
    if (this.codebaseForm.invalid) return;

    const payload = {
      codebase_name: this.codebaseForm.value.codeBaseName,
      github_url: this.codebaseForm.value.gitHubURL,
      username: this.codebaseForm.value.username,
      token: this.codebaseForm.value.token,
      branch: this.codebaseForm.value.branch,
      folder_path: this.codebaseForm.value.folderPath
    };
    console.log('Submitting:', payload);
    this.dashBoardService.extractCode(payload).subscribe({
      next: () => {
        this.toastr.success('Code Base Syncing  Stared', 'Success');
        this.syncYourCodeBaseClicked = false;
        this.getCodeBases();
      },
      error: (error) => {
        console.log(error);
        this.toastr.error('CodeBase Syncing Failed ...', error.error.detail);
        this.syncYourCodeBaseClicked = false;
      },
    })
  }
  navigateToCodeBaseChat(codeBaseId: string, codeBaseName: string) {
    this.router.navigate(['/codebase-chat'], {
      queryParams: {
        codeBaseId: codeBaseId,
        codeBaseName: codeBaseName
      }
    });
  }

  deleteCodeBaseContext(codeBaseId: string) {
    this.dashBoardService.DeleteCodeBasdeContext(codeBaseId).subscribe({
      next: () => {
        this.toastr.success(`Code Base Deleted entirely  ...`)
      },
      error: (error) => {
        console.log(error);
        this.toastr.error('Error in deleting codeBase .');
      }
    })
  }

}
