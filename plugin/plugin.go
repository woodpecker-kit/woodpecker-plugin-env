package plugin

import (
	"fmt"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"os"
	"strconv"
	"strings"
)

type (
	// Plugin plugin all config
	Plugin struct {
		Name           string
		Version        string
		WoodpeckerInfo *wd_info.WoodpeckerInfo
		Config         Config
	}
)

func (p *Plugin) Exec() error {
	// replace this code with your plugin implementation

	var sb strings.Builder
	_, _ = fmt.Fprint(&sb, "-> just print basic env:\n")
	paddingMax := strconv.Itoa(p.Config.PaddingLeftMax)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCiWorkflowName, p.WoodpeckerInfo.CurrentInfo.CurrentWorkflowInfo.CiWorkflowName)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyWoodpeckerBackend, p.WoodpeckerInfo.CiSystemInfo.WoodpeckerBackend)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiMachine, p.WoodpeckerInfo.CiSystemInfo.CiMachine)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiSystemPlatform, p.WoodpeckerInfo.CiSystemInfo.CiSystemPlatform)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyRepositoryCiName, p.WoodpeckerInfo.RepositoryInfo.CIRepoName)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyRepositoryCiOwner, p.WoodpeckerInfo.RepositoryInfo.CIRepoOwner)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitBranch, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitBranch)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitRef, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitRef)

	appendStrBuilderNewLine(&sb)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineNumber, p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineNumber)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineEvent, p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent)

	switch p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent {
	case wd_info.EventPipelineTag:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitTag, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTag)
	case wd_info.EventPipelinePullRequest:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequest, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequestLabels, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitSourceBranch, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitTargetBranch, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch)
	case wd_info.EventPipelinePullRequestClose:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequest, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequestLabels, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitSourceBranch, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitTargetBranch, p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch)
	case wd_info.EventPipelineRelease:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPreRelease, strconv.FormatBool(p.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPreRelease))
	}

	appendStrBuilderNewLine(&sb)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineUrl, p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineForgeUrl, p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineForgeUrl)

	appendStrBuilderNewLine(&sb)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiCommitBranch, p.WoodpeckerInfo.PreviousInfo.PreviousCommitInfo.CiPreviousCommitBranch)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiCommitRef, p.WoodpeckerInfo.PreviousInfo.PreviousCommitInfo.CiPreviousCommitRef)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiPipelineEvent, p.WoodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineEvent)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiPipelineStatus, p.WoodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineStatus)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiPipelineUrl, p.WoodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineUrl)

	if len(p.Config.EnvPrintKeys) > 0 {
		appendStrBuilderNewLine(&sb)
		_, _ = fmt.Fprint(&sb, "-> start print keys env:\n")
		for _, key := range p.Config.EnvPrintKeys {
			appendEnvStrBuilder(&sb, paddingMax, key, os.Getenv(key))
		}
		_, _ = fmt.Fprint(&sb, "-> end print keys env\n")
	}

	wd_log.Verbosef("%s", sb.String())
	return nil
}

func appendStrBuilderNewLine(sb *strings.Builder) {
	_, _ = fmt.Fprintf(sb, "\n")
}

func appendEnvStrBuilder(sb *strings.Builder, paddingMax string, key string, value string) {
	_, _ = fmt.Fprintf(sb, "%-"+paddingMax+"s %s\n", key, value)
}
